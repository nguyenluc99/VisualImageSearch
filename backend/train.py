from pathlib import Path
import os

import numpy as np
from PIL import Image

from feature_extractor import FeatureExtractor

if __name__ == '__main__':
    fe = FeatureExtractor()

    arr = None
    batch_size = 128
    count = 0
    batch_count = 0

    # TODO: sort by name
    for idx, img_path in enumerate(sorted(Path("./images").glob("*.jpg"), key=os.path.getmtime)[:500]):
        feature = fe.extract(img=Image.open(img_path))
        feature_path = Path("./feature/all") / (img_path.stem + ".npy")  # e.g., ./static/feature/xxx.npy

        # save to feature/all
        np.save(feature_path, feature)

        # save to feature/batch
        if arr is not None:
            arr = np.append(arr, feature, axis=0)
            count += 1 
        else:
            arr = feature
            count = 1

        if count == batch_size:
            np.save(f"feature/batch/{batch_count}.npy", arr)
            batch_count += 1
            arr = None
            count = 0
        # save the order of image in feature/batch
        with open("test.txt", "a") as index_file:
            index_file.write(f"{idx} - {img_path.stem}\n")


    # TODO: train autofaiss model
