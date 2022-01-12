import React from "react";
import { Container, Row, Col } from "react-bootstrap";

export default function ListImg(props) {
    function gcd(a, b) {
        while (a != b) {
            if (a > b) a = a - b;
            else b = b - a;
        }
        return a;
    }
    console.log(props.data);

    return (
        <>
            <div className="list_img">
                <Container>
                    <h1 className="title">
                        Thuật toán <br /> xxx
                    </h1>
                    <Row>
                        {props.data.map((item, index) => {
                            console.log(item, )
                            return (
                                <Col key={index} xs={6} lg={3} md={4}>
                                    <img className="img_sp" src={`http://localhost:5000/images/${item.image}`} alt="" />
                                    <span>{item.name}</span>
                                </Col>
                            );
                        })}
                    </Row>
                </Container>
            </div>
        </>
    );
}
