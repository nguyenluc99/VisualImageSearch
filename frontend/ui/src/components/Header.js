import React, { useState } from "react";
import { Container } from "react-bootstrap";

export default function Header() {
    const [imgUpload, setImgUpload] = useState("");
    const [widthImg, setWidthImg] = useState(500);
    const [heightImg, setHeightImg] = useState(300);
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
        reader.readAsDataURL(file);
    }

    return (
        <div className="header">
            <Container>
                <h1 className="title">Computer Vision Application</h1>

                <div className="card" style={{ width: widthImg, height: heightImg }}>
                    <span className="btn_upload">
                        <i className="fas fa-image"></i>
                    </span>

                    <img src={imgUpload} style={{ display: imgUpload ? "" : "none", width: "100%", height: "100%", position: "absolute", transition: "0.3s" }} />

                    <input id="upload_file" type="file" onChange={uploadImg} title="upload new img" />
                </div>
                <div className="btn_submit">
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
