from PIL import Image
from datetime import datetime
from flask import Flask, request, render_template, send_from_directory, jsonify
from image_search import autofaiss_img_search

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
        scores = autofaiss_img_search(img)
        print(scores)

        # return render_template('index.html',
        #                        query_path=uploaded_img_path,
        #                        scores=scores)
        # response json template
        return jsonify({
            "query_path": uploaded_img_path,
            "scores": scores
        })
    else:
        return render_template('index.html')


@app.route('/images/<path:path>')
def image(path):
    return send_from_directory('images', path)

if __name__=="__main__":
    app.run("0.0.0.0", debug=True)
