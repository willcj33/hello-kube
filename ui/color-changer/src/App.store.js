import { createStore } from 'redux'
import colorGrid from './App.reducers'
let store = createStore(colorGrid);
export default store;