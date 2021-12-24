import React, { useEffect, useState } from "react";
import Header from "./components/Header";
import ListImg from "./components/ListImg";

const BACKEND_URL = "https://freeapi.code4func.com/api/v1";

function getImage(src) {
    let img = new Image();
    img.src = src;
    return img;
}

function App() {
    const [data, setData] = useState([]);
    useEffect(() => {
        (async () => {
            fetch(`${BACKEND_URL}/cate/list`)
                .then((response) => response.json())
                .then(
                    (success) => {
                        console.log({ success });
                        (async () => {
                            setData(
                                success.data.map((img) => {
                                    return {
                                        src: img.cateImage,
                                        caption: img.cateName,
                                    };
                                })
                            );
                        })();
                    },
                    (error) => {
                        console.error(error);
                    }
                );
        })();

        return () => {};
    }, []);

    return (
        <div className="App">
            <Header />
            <ListImg data={data} />
        </div>
    );
}

export default App;
