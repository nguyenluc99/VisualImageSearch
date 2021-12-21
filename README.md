# VisualImageSearch
Visual image search project at our Computer Vision course in Hanoi University of Science and Technology

# Instructions

## Back-end
Copy the images folder to the backend folder

Create folder `backend/feature`. This folder stores feature of image in `npy` format.

Run `pip3 install -r requirements.txt` to install needed library.

Run `python3 train.py` to extract feature of image.

Then, run `python3 api.py` to serve the api. (not working right now)
 
## Quick test with command line:
Run `python3 test.py image_path` to test the image specified by the `image_path`.

