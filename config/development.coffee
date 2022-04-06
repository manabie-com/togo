mongoCfg =
  user: 'visitor'
  password: 'xrrG8qNdl9McOEpF'
  database: 'manabie_togo'


module.exports = {
  mongo: mongoCfg
  app:
    port: 8004
  mongoConnectionString: "mongodb+srv://#{mongoCfg.user}:#{mongoCfg.password}@togo.l0zat.mongodb.net/#{mongoCfg.database}?retryWrites=true&w=majority"
  prefixCollection: 'mnb'
  prefixModel: 'mnb'
}

