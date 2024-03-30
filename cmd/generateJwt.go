package cmd

import (
	Helpers "cli-go/internal/featureArch/helpers"
	"fmt"
)

func GenerateJwtMiddleware() {
	code := `import jwt from 'jsonwebtoken';
import { Request, Response, NextFunction } from 'express';
import HttpException from '../exceptions/httpException';
import { config } from '../config';

require('dotenv').config();


export default class AuthMiddleware{
  static secretKey: string = config.JWT_TOKEN_SECRET;
static verifyTokenMiddleware = (
  req: Request,
  res: Response,
  next: NextFunction,
  // eslint-disable-next-line consistent-return
): void | HttpException => {
  const token = req.headers.authorization?.split(' ')[1];

  if (!token) {
    return next(
      new HttpException(401, JSON.stringify({ error: 'token not provided' })),
    );
  }

  try {
    jwt.verify(token, secretKey);
    next();
  } catch (error) {
    return next(
      new HttpException(403, JSON.stringify({ error: 'invalid token' })),
    );
  }
};
static generateToken=(payload: any): string {
    // Generate JWT token with payload and secret key
    const token = jwt.sign(payload, secretKey);
    return token;
}

}
`
	path := "src/middleware/auth.ts"
	Helpers.GenerateJavascriptFile(path, code)
	fmt.Printf("✅ JWT middleware generated at %s", path)
}
