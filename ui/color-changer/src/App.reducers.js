const initialState = {
    gridSize: {
        width: 3,
        height: 3
    },
    grid: [
        [{id: '00', color: '#000'},{id: '01', color: '#000'},{id: '02', color: '#000'}],
        [{id: '10', color: '#000'},{id: '11', color: '#000'},{id: '12', color: '#000'}],
        [{id: '20', color: '#000'},{id: '21', color: '#000'},{id: '22', color: '#000'}]
    ]
};

export default (state = initialState, action) => {
    switch (action.type) {
        case 'CHANGE_COLOR':
            return Object.assign({}, state, { 
                grid: state.grid.map(r => r.map(c => {
                    if(c.id === action.id) {
                        return Object.assign({}, c, { color: action.color });
                    }
                    return c;
                })) 
            });
        case 'UPDATE_GRID':
            let newGrid = [];
            for(let h = 0; h < action.height; h++) {
                newGrid[h] = state.grid[h] && state.grid[h].length ? state.grid[h] : [];
                for(let w = 0; w < action.width; w++) {
                    newGrid[h][w] = newGrid[h][w] && newGrid[h][w].color ? newGrid[h][w] : { id: `${h}${w}`, color: '#000' };
                }
            }
            return {
                gridSize: {
                    width: action.width,
                    height: action.height
                },
                grid: newGrid
            };
        default:
            return state;
    }
};