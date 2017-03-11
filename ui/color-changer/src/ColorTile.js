import React, { Component } from 'react';

export default class ColorTile extends Component {
    constructor (props) {
        super(props);
    }
    render () {
        return (
            <div key={this.props.id} id={this.props.id} className="color-tile" style={{backgroundColor: this.props.color}}></div>
        );
    }
}