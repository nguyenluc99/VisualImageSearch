# Visual Image Search
Visual image search project at our Computer Vision course in Hanoi University of Science and Technology

# Instructions

## Crawl data.
cd to `crawler` folder, then run `go main.go` to crawl data.
The crawled images is store in the `images` folder.

## Back-end
Structure of backend.
```
├── api.py
├── config
│   └── tiki.yml
├── feature
│   ├── batch # contains .npy matrix (output of train.py file)
│   ├── pyretri
├── feature_extractor.py # implements a class, that help we extract vector from image.
├── idx2path.pkl # a python dictionary, that is a mapping from index to image path
├── images # crawled images.
├── image_search.py # contains function that call the search model.
├── models
│   ├── index_infos.json
│   └── knn.index
├── mongo.py # create connection to the MongoDB database
├── pyretri # pyretri project
│   ├── config # contain config file
│   ├── datasets # sample dataset of pyretri
│   ├── evaluate # contains code that create the evaluation for model
│   ├── extract # extract vector from image.
│   ├── index # evaluate model and find similar image.
│   ├── __init__.py
│   ├── models
│   ├── __pycache__
│   ├── utils
│   └── version.py
├── requirements.txt # contains needed libraries
├── static
│   └── uploaded # contains uploaded images.
├── test.txt
└── train.py
```

Copy the images folder to the backend folder

Create folder `backend/feature`. This folder stores feature of image in `npy` format.

Run `pip3 install -r requirements.txt` to install needed library.

Download models.
- Download `idx2path.pkl` file from this [link](https://drive.google.com/file/d/14zyVie6A2-VgHefPaSE8MBCcnY93lX5T/view?usp=sharing) into `backend` directory.
- Download and unzip autofaiss model from [here](https://drive.google.com/drive/folders/136jsKFdcfj8Q-IroGFusjFv3tbpcpmd1?usp=sharing), into `backend/models` directory
- Download and unzip pyretri models from [here](https://drive.google.com/file/d/1kVDvPIrhFIYZJjJC2vXiiaimQtyyLSvs/view?usp=sharing), into `backend/feature` directory.

Then, run `python3 api.py` to serve the api, then, cd to `frontend/ui` directory and run `npm start` to run the website.


# Main code snippets of the project

- Search similar images.
![image](https://user-images.githubusercontent.com/36019052/149182248-8ef72b65-1f25-4794-9535-2934b0fe6e3f.png)

- From similar images, get at most 10 products.
![image](https://user-images.githubusercontent.com/36019052/149182345-5a06569c-d64a-4355-a37d-a2ee213c5497.png)

- Feature Extractor class
![image](https://user-images.githubusercontent.com/36019052/149182531-172a34d7-c47d-4a85-a9f2-ff5c732304f1.png)

- Search API
![image](https://user-images.githubusercontent.com/36019052/149182904-3f680a7e-089a-45db-b380-bb6c2b823100.png)

- Crawl product list
![image](https://user-images.githubusercontent.com/36019052/149183029-5ca6d8a7-e781-4170-9738-f418190dd1c8.png)

- Get product detail and save
![image](https://user-images.githubusercontent.com/36019052/149183133-14522495-bdc4-428e-9026-ac8ac9040d25.png)

- Download image
![image](https://user-images.githubusercontent.com/36019052/149183164-bd43e3a6-1c80-4d59-81d8-4477dda00db7.png)

