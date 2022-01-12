import React, { useState } from "react";
import { Container } from "react-bootstrap";

export default function Header(props) {
    const setAutofaissList = props.setAutofaissList;
    const [imgUpload, setImgUpload] = useState("");
    const [file, setFile] = useState(null);
    const [widthImg, setWidthImg] = useState(500);
    const [heightImg, setHeightImg] = useState(300);
    const formData = new FormData();
  
    function uploadImg(e) {
        const file = e.target.files[0];
        var reader = new FileReader();
        reader.onload = function (ee) {
            var img = new Image();
            img.onload = function () {
                if (this.width > this.height) {
                    if (this.width > 500) {
                        setWidthImg(500);
                        setHeightImg((500 * this.height) / this.width);
                    } else {
                        setWidthImg(this.width);
                        setHeightImg(this.height);
                    }
                } else {
                    if (this.height > 500) {
                        setHeightImg(500);
                        setWidthImg((500 * this.width) / this.height);
                    } else {
                        setWidthImg(this.width);
                        setHeightImg(this.height);
                    }
                }
            };
            img.src = ee.target.result;
            setImgUpload(ee.target.result);
        };
        // if (formData.has("file")) {
        //     formData.delete("file");
        // }
        console.log("load file", file)
        formData.append("file", e.target.files[0]);
        setFile(e.target.files[0]);

        reader.readAsDataURL(file);
    }
    
    const search = () => {
      console.log(formData)
        formData.append("model_name", "autofaiss");
        formData.append("file", file);
        fetch("http://localhost:5000/search", {
            method: "POST",
            body: formData
        }).then(function (res) {
          return res.json()
        }).then(data => {
          console.log("data", data)
          setAutofaissList(data.products)
        });
    }

    return (
        <div className="header">
            <Container>
                <div className="card" style={{ width: widthImg, height: heightImg }}>
                    <span className="btn_upload">
                        <i className="fas fa-image"></i>
                    </span>

                    <img src={imgUpload} style={{ display: imgUpload ? "" : "none", width: "100%", height: "100%", position: "absolute", transition: "0.3s" }} />

                    <input id="upload_file" type="file" onChange={uploadImg} title="upload new img" />
                </div>
                <div className="btn_submit" onClick={() => search()}>
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                    Submit
                </div>
            </Container>
        </div>
    );
}
