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
                            return (
                                <Col key={item.src + index} xs={6} lg={3} md={4}>
                                    <img className="img_sp" src={item.src} alt="" />
                                </Col>
                            );
                        })}
                    </Row>
                </Container>
            </div>
            <div className="list_img">
                <Container>
                    <h1 className="title">
                        Thuật toán <br /> xxx
                    </h1>
                    <Row>
                        {props.data.map((item, index) => {
                            return (
                                <Col key={item.src + index} xs={6} lg={3} md={4}>
                                    <img className="img_sp" src={item.src} alt="" />
                                </Col>
                            );
                        })}
                    </Row>
                </Container>
            </div>
        </>
    );
}
