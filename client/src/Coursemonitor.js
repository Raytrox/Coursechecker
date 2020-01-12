import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Button } from "semantic-ui-react";

let endpoint = "http://localhost:8010";

class CourseMonitor extends Component {

  constructor(props) {
    super(props);
    this.state = {
        data: ''
    };
    this.click = this.click.bind(this);
  }

  click() {
    axios.post(endpoint + "/monitor", {
      dept: document.getElementById("dept").value,
      number: document.getElementById("number").value,
      section: document.getElementById("section").value,
      receiver: document.getElementById("email").value,
      status: "monitoring"
    })
      .then((res) => {
        console.log(res.data)
        window.alert("Monitoring has started, you'll receive an email if there are seats available!")
      })
      .catch((err) => {
        console.log(err)
        window.alert("Please double check the information entered!")
      });
  }

  getInfo() {

  }

  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h1">
            Course Monitor
          </Header>
        </div>
        <div class="ui section divider"></div>
        <div class="ui labeled input">
          <div class="ui label">
          dept
          </div>
          <input type="text" placeholder="" id="dept" />
        </div>
        <div class="ui labeled input">
          <div class="ui label">
          number
          </div>
          <input type="text" placeholder="" id="number"/>
        </div>
        <div class="ui labeled input">
          <div class="ui label">
          section
          </div>
          <input type="text" placeholder="" id="section"/>
        </div>
        <div class="ui labeled input">
          <div class="ui label">
          email address
          </div>
          <input type="text" placeholder="" id="email"/>
        </div>
        <div class="ui section divider"></div>
        <button class="ui primary button" onClick = {this.click}>
        Start monitoring
        </button>
       <button class="ui button" aria-label="check all monitoring info" onClick = {this.getInfo}>
         check all monitoring info
       </button>
      </div>
    );
  }
}

export default CourseMonitor;