export function changeColor(id, color) { 
    return { 
        type: 'CHANGE_COLOR', 
        id, 
        color 
    } 
};

export function updateGridSize(width, height) { 
    return { 
        type: 'UPDATE_GRID', 
        width, 
        height 
    } 
};
