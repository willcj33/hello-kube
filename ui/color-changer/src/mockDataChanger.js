import Store from './App.store'
import { changeColor, updateGridSize } from './App.actions'

const getRandomColor = () => {
    let r = Math.floor(Math.random() * 256).toString(16);
    let g = Math.floor(Math.random() * 256).toString(16);
    let b = Math.floor(Math.random() * 256).toString(16);
    return `#${r + g + b}`;
};

const getRandomTileId = (width, height) => {
    let wi = Math.floor(Math.random() * width);
    let hi = Math.floor(Math.random() * height);
    return `${hi}${wi}`;
};

setInterval(() => {
    let gridSize = Store.getState().gridSize;
    let id = getRandomTileId(gridSize.width, gridSize.height);
    let color = getRandomColor();
    Store.dispatch(changeColor(id, color));
}, 1000);

setInterval(() => {
    let w = Math.ceil(Math.random() * 6);
    let h = Math.ceil(Math.random() * 6);
    Store.dispatch(updateGridSize(w, h));
}, 10000)