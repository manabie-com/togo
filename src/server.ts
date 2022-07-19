import express, { Express, Request, Response } from 'express';
import bodyParser from 'body-parser';
import dotenv from 'dotenv';

dotenv.config();
const app: Express = express();
const port = process.env.PORT || 3009;

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());


app.get('/', (req: Request, res: Response) => {
	res.send('Welcome Togo');
})

app.listen(port, () => {
	console.log(`Server running on port: ${port}`);
})
