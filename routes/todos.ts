import express, { Request, Response } from "express";

let router = express.Router();

router.route("/").post(async (req: Request, res: Response) => {
  try {
    const params = req.body;

		return res.sendStatus(200)
  } catch (e) {
    console.log(`Error: ${e}`);
    return res.sendStatus(500);
  }
});

export default router;