import Store from './App.store'
import { changeColor } from './App.actions'

const numberToCoordinate = (theNumber, gridSize) => {
    theNumber += 1;
    let x = Math.ceil(theNumber / gridSize.width) - 1;
    let y = (theNumber % gridSize.width);
    return `${x}${y}`;
};

let conn = new WebSocket("ws://192.168.99.100:31384/ws?interval=2000&count=9");

conn.onmessage = function (evt) {
    let gridSize = Store.getState().gridSize;
    var data = JSON.parse(evt.data);
    let id = numberToCoordinate(data.which, gridSize);
    Store.dispatch(changeColor(id, `#${data.color}`));
};

conn.onclose = function (evt) {
    console.log("closed");
};

