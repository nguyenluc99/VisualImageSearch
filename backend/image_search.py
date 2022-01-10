import os
from pathlib import Path

import faiss
import numpy as np

from PIL import Image

from pyretri.config import get_defaults_cfg, setup_cfg
from pyretri.datasets import build_transformers
from pyretri.models import build_model
from pyretri.extract import build_extract_helper
from pyretri.index import build_index_helper, feature_loader
from pyretri.evaluate import build_evaluate_helper

from feature_extractor import FeatureExtractor
import pickle

# Read image features
fe = FeatureExtractor()
features = []
img_paths = []
autofaiss_img_paths = []


# for feature_path in sorted(Path("./images").glob("*"), key=os.path.getmtime):
#     autofaiss_img_paths.append(str(Path("./images") / (feature_path.stem + ".jpg")))

# for feature_path in sorted(Path("./feature/all").glob("*.npy"), key=os.path.getmtime):
#     features.append(np.load(feature_path))
#     img_paths.append(Path("./images") / (feature_path.stem + ".jpg"))
# features = np.array(features)

autofaiss_model = faiss.read_index('models/knn.index')

autofaiss_idx2path = {}
with open("idx2path.pkl", "rb") as f:
    autofaiss_idx2path = pickle.load(f)

def autofaiss_img_search(img_path):
    img = Image.open(img_path)  # PIL image
    query_vector = fe.extract(img)
    distances, indices = autofaiss_model.search(query_vector.reshape(1, -1), 50)
    print(distances, indices)
    result_paths = []
    for idx in indices[0]:
        result_paths.append(str(autofaiss_idx2path[idx]))
    # print(result_paths)
    return result_paths

cfg = get_defaults_cfg()
cfg = setup_cfg(cfg, './config/pyretri.yml', [])
model = build_model(cfg.model)

transformers = build_transformers(cfg.datasets.transformers)

# # load gallery features
gallery_fea_out, gallery_info, _ = feature_loader.load(cfg.index.gallery_fea_dir, cfg.index.feature_names)
def pyretrive_img_search(img_path):
    img = Image.open(img_path)
    img_tensor = transformers(img)
  
    extract_helper = build_extract_helper(model, cfg.extract)
    img_fea_info = extract_helper.do_single_extract(img_tensor)
    stacked_feature = list()
    for name in cfg.index.feature_names:
        assert name in img_fea_info[0], "invalid feature name: {} not in {}!".format(name, img_fea_info[0].keys())
        stacked_feature.append(img_fea_info[0][name].cpu())
    img_fea = np.concatenate(stacked_feature, axis=1)


    # build helper and single index feature
    index_helper = build_index_helper(cfg.index)
    index_result_info, query_fea, gallery_fea = index_helper.do_index(img_fea, img_fea_info, gallery_fea_out)
  
    index_result_info = index_result_info[0]
    # return top 50 results
    result_paths = []
    for i in index_result_info["ranked_neighbors_idx"][:50]:
        path = gallery_info[i]["path"]
        # name = path.split("/")[-1]
        result_paths.append(path)
    return result_paths


def image_paths_to_product_list(image_paths):
    """
    From list of product images, get 10 products from this list
    Note: one product may have more than 1 image
    """
    product_list = []
    for image_path in image_paths:
        name = image_path.split("/")[-1]
        product_id = name.split("-")[0]
        image_path = "/".join(image_path.split("/")[1:])
        if product_id not in product_list:
            product_list.append({"product_id": product_id, "image": image_path})
    return product_list[:10]

if __name__ == "__main__":
    # paths = pyretrive_img_search("images/chao/54625005-0.jpg")
    # print(image_paths_to_product_list(paths))
    paths = autofaiss_img_search("images/trai-cay/117473293-1.jpg")
    for i in paths:
        print(i)
