from datetime import datetime

from flask import Flask, jsonify, render_template, request, send_from_directory
from PIL import Image
from flask_cors import CORS

from image_search import (autofaiss_img_search, image_paths_to_product_list,
                          pyretrive_img_search)
from mongo import get_database

app = Flask(__name__)
CORS(app)
@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        file = request.files['file']

        # Save query image
        img = Image.open(file.stream)  # PIL image
        uploaded_img_path = "static/uploaded/" + datetime.now().isoformat().replace(":", ".") + "_" + file.filename
        img.save(uploaded_img_path)

        # Run search
        scores = autofaiss_img_search(uploaded_img_path)
        images = image_paths_to_product_list(ret)

        # return render_template('index.html',
        #                        query_path=uploaded_img_path,
        #                        scores=scores)
        # response json template
        return jsonify({
            "product": product,
        })
    else:
        return render_template('index.html')

db = get_database()

@app.route('/search', methods=["POST"])
def search():
    file = request.files['file']
    model_name = request.form['model_name']
    
    # Save query image
    img = Image.open(file.stream)  # PIL image
    uploaded_img_path = "static/uploaded/" + datetime.now().isoformat().replace(":", ".") + "_" + file.filename
    img = img.convert('RGB')
    img.save(uploaded_img_path)

    products = []
    if model_name == 'pyretri':

        # Run search
        ret = pyretrive_img_search(uploaded_img_path)
        products = image_paths_to_product_list(ret)
    elif model_name == 'autofaiss':
        ret = autofaiss_img_search(uploaded_img_path)
        products = image_paths_to_product_list(ret)
    product_ids = [int(product['product_id']) for product in products]
    print(product_ids)
    ret = db.products.find({"id": {"$in": product_ids}}, {"_id": 0, "id": 1, "name": 1, "price": 1, "urlkey": 1})
    ret = list(ret)
    for idx, product in enumerate(ret):
        product['image'] = products[idx]['image']
    return jsonify({
        "products2": products,
        "products": ret,
    })

@app.route('/images/<path:path>')
def image(path):
    return send_from_directory('images', path)

if __name__=="__main__":
    app.run("0.0.0.0", debug=False)
