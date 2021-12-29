import os
from pathlib import Path

import faiss
import numpy as np

from feature_extractor import FeatureExtractor

# Read image features
fe = FeatureExtractor()
features = []
img_paths = []
autofaiss_img_paths = []


for feature_path in sorted(Path("./images").glob("*"), key=os.path.getmtime):
    autofaiss_img_paths.append(str(Path("./images") / (feature_path.stem + ".jpg")))

for feature_path in sorted(Path("./feature/all").glob("*.npy"), key=os.path.getmtime):
    features.append(np.load(feature_path))
    img_paths.append(Path("./images") / (feature_path.stem + ".jpg"))
features = np.array(features)

model = faiss.read_index('models/knn.index')

def autofaiss_img_search(img):
    query_vector = fe.extract(img)
    distances, indices = model.search(query_vector.reshape(1, -1), 10)
    return [(str(distances[0][i]), autofaiss_img_paths[indices[0][i]]) for i in range(10)]
