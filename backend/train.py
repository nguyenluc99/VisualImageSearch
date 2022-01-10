from pathlib import Path
import os

import numpy as np
from PIL import Image, ImageFile
ImageFile.LOAD_TRUNCATED_IMAGES = True

from feature_extractor import FeatureExtractor
import pickle

if __name__ == '__main__':
    fe = FeatureExtractor()

    arr = None
    batch_size = 128
    count = 0
    batch_count = 0
    idx2path = {}

    # TODO: sort by name
    # for idx, img_path in enumerate(sorted(Path("./images").glob("*.jpg"), key=os.path.getmtime)[:500]):
    for root, dirs, files in os.walk("./images"):
        for dir in dirs:
            for img_path in sorted(Path(root).glob(dir + "/*.jpg"), key=os.path.getmtime)[:15]:
                idx2path[count] = img_path
                feature = fe.extract(img=Image.open(img_path))
                feature_path = Path("./feature/all") / (img_path.stem + ".npy")  # e.g., ./static/feature/xxx.npy

                # save to feature/all
                np.save(feature_path, feature) # for normal search

                # save to feature/batch
                if arr is not None:
                    arr = np.append(arr, feature, axis=0)
                    count += 1
                else:
                    arr = feature
                    count = 0

                idx2path[batch_size * batch_count + count] = img_path
                print(batch_size * batch_count + count, img_path)
                if count + 1 == batch_size:
                    np.save(f"feature/batch/{batch_count}.npy", arr)
                    batch_count += 1
                    arr = None
                    count = 0

    np.save(f"feature/batch/{batch_count}.npy", arr)
    pickle.dump(idx2path, open("idx2path.pkl", "wb"))
