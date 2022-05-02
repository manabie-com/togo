print('Init DB ###############################################');
db = db.getSiblingDB('togo-db');
db.createUser({
  user: 'local-togo',
  pwd: 'root',
  roles: [
    {
      role: 'readWrite',
      db: 'togo-db'
    }
  ]
});
