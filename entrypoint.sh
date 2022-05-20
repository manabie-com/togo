sleep 8 #Sleep to wait db ready
npm run build
npm run migrate:up
npm run seed:run
npm run start:prod