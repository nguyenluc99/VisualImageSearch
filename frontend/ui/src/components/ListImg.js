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
    if (!props.data?.products) return <div />;

    return (
        <>
            <div className="list_img">
                <Container>
                    <h1 className="title">
                        {props.name}
                    </h1>
      <span>{props.data.elapsed_time}ms</span>
                    <Row>
                        {props.data.products.map((item, index) => {
                            console.log(item, )
                            return (
                                <Col key={index} xs={6} lg={2} md={4}>
                                  <div style={{display: 'flex', flexDirection: 'column'}} >
                                    <img className="img_sp" src={`http://localhost:5000/images/${item.image}`} alt="" style={{height: 150, width: 'auto'}} />
                                    <a href={`https://tiki.vn/${item.urlkey}.html`}>{item.name}</a>
                                  </div>
                                </Col>
                            );
                        })}
                    </Row>
                </Container>
            </div>
        </>
    );
}
