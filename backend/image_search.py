from feature_extractor import FeatureExtractor
from pathlib import Path
import numpy as np

# Read image features
fe = FeatureExtractor()
features = []
img_paths = []
for feature_path in Path("./feature").glob("*.npy"):
    features.append(np.load(feature_path))
    img_paths.append(Path("./images") / (feature_path.stem + ".jpg"))
features = np.array(features)


def img_search(img):
    query = fe.extract(img)
    dists = np.linalg.norm(features - query, axis=1)  # L2 distances to features
    ids = np.argsort(dists)[:30]  # Top 30 results
    scores = [(dists[id], img_paths[id]) for id in ids]
    print(ids, scores)
    return scores

