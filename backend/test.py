import sys
from PIL import Image
from image_search import img_search


def main(filename):
    img = Image.open(filename)
    scores = img_search(img)
    print(len(scores))
    print(scores)

if __name__ == '__main__' :
    if len(sys.argv) < 2: 
        print("Input a image test file: python3 test.py <filename>");
        exit(1)
    main(sys.argv[1])

