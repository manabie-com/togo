const knex = require('../db/knex');

module.exports.getAllUsers = async (req, res) => {
    const users = await knex('users').select('*');
    res.status(200).json({ users });
}

module.exports.createUser = async (req, res) => {
    const results = await knex('users').insert(req.body);
    res.status(201).json({ id: results[0] });
}

module.exports.updateUser = async (req, res) => {
    const id = await knex('users').where('id', req.params.id).update(req.body);
    res.status(200).json({ id });
}

module.exports.deleteUser = async (req, res) => {
    await knex('users').where('id', req.params.id).del();
    res.status(200).json({ success: true });
}