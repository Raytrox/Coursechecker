import React from "react";
import "./App.css";

// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";

// import the ToDoList component
import CourseMonitor from "./CourseMonitor";
function App() {
  return (
    <div>
      <Container>
        <CourseMonitor />
      </Container>
    </div>
  );
}

export default App;
