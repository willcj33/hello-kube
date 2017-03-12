import React, { Component } from 'react';
import Store from './App.store'
import ColorTile from './ColorTile'
import './App.css';

class App extends Component {
  constructor (props) {
    super(props);
    this.state = Store.getState();
  }

  componentDidMount() {
    this.unsubscribe = Store.subscribe(() => this.setState(Store.getState()));
  }

  componentWillUnmount() {
    this.unsubscribe();
  }

  render() {
    return (
      <div className="container">
        {this.state.grid.map((r, index) => <div key={index} className="row">
          {r.map(c => <ColorTile key={c.id} id={c.id} color={c.color}></ColorTile>)}
        </div>)}
      </div>
    );
  }
}

export default App;