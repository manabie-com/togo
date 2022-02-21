let users = [
    {'id':1, 'name': 'A', metadata: {max_tasks: 2}},
    {'id':2, 'name': 'B', metadata: {max_tasks: 3}}
];

module.exports.findById = (id) => {
    return users.find(user => user.id === id); 
};
