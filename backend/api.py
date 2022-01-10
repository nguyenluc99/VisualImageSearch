from PIL import Image
from datetime import datetime
from flask import Flask, request, render_template, send_from_directory, jsonify
from image_search import autofaiss_img_search, pyretrive_img_search, image_paths_to_product_list

app = Flask(__name__)

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

    return jsonify({
        "products": products
    })

@app.route('/images/<path:path>')
def image(path):
    return send_from_directory('images', path)

if __name__=="__main__":
    app.run("0.0.0.0", debug=False)
