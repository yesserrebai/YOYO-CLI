package innerStructure

import (
	Helpers "cli-go/internal/featureArch/helpers"
	"fmt"
)

func CreateFolders(projectName string, folders []string) {
	for _, folder := range folders {
		Helpers.CreateFolder(fmt.Sprintf("%s/%s", projectName, folder))
	}

}
func CreateEnvFilesJSON(projectName string, environments []string) {
	jsonData := `{SUPPRESS_NO_CONFIG_WARNING: true}`
	for _, env := range environments {
		Helpers.GenerateJsonFile(projectName+"/src/config/"+env+".json", jsonData)
	}
}
func CreateConfigFile(projectName string) {
	jsonData := `import { DotenvParseOutput } from 'dotenv';
// eslint-disable-next-line import/no-unresolved
import 'reflect-metadata';

import { validateConfig } from './env.validation';

const parsed = process.env;

const config = validateConfig(parsed as DotenvParseOutput);
export default config;
`
	path := projectName + "/src/config/env/config.ts"
	Helpers.GenerateJavascriptFile(path, jsonData)
}
func CreateEnvValidationFile(projectName string) {
	jsonData := `import { plainToInstance } from 'class-transformer';
import { validateSync } from 'class-validator';
import { EnvironmentVariables } from './environment-variables';

export function validateConfig(
  config: Record<string, unknown>,
): EnvironmentVariables {
  const validatedConfig = plainToInstance(EnvironmentVariables, config, {
    enableImplicitConversion: true,
    exposeDefaultValues: true,
  });
  let errors: any;
  if (process.env.NODE_ENV == 'test') {
    errors = validateSync(validatedConfig, {
      skipMissingProperties: true,
    });
  } else {
    errors = validateSync(validatedConfig, {
      skipMissingProperties: false,
    });
  }

  if (errors.length > 0) {
    const errMessages: string[] = [];

    errors.forEach((e: { constraints: { [type: string]: string } }) => {
      const errConstraints = e.constraints as { [type: string]: string };

      // eslint-disable-next-line no-restricted-syntax
      for (const [, errMessage] of Object.entries(errConstraints)) {
        errMessages.push(errMessage);
      }
    });

    throw new Error(
      Config initialization error: \n${errMessages.join(',\n')},
    );
  }

  return validatedConfig;
}
`
	path := projectName + "/src/config/env/env-validations.ts"
	Helpers.GenerateJavascriptFile(path, jsonData)
}
func CreateEnvironmentVariablesFile(projectName string) {
	jsonData := `import { IsNotEmpty, IsNumber, IsString, IsEnum } from 'class-validator';

import { Environment } from '../enum';

export class EnvironmentVariables {
  @IsEnum(Environment)
  public NODE_ENV: string;

  @IsNumber()
  @IsNotEmpty()
  public PORT: number;
}
`
	path := projectName + "/src/config/env/environment-variables.ts"
	Helpers.GenerateJavascriptFile(path, jsonData)
}
func CreateEnumFile(projectName string) {
	code := `export enum Environment {
  DEV = 'dev',
  STAGING = 'staging',
  PRE_PROD = 'preprod',
  PRODUCTION = 'prod',
  TEST = 'test',
}`
	path := projectName + "/src/config/enum.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateIndexFile(projectName string) {
	code := `import config from './env/config';

export { config };
`
	path := projectName + "/src/config/index.ts"
	Helpers.GenerateJavascriptFile(path, code)

}
func CreateServerFile(projectName string) {
	code := `import app from './app';
require('dotenv').config();
const port: number = Number(process.env.PORT) || 8084;
app.listen(port, () => {
  console.info(
    Listening on port ${port}, environment:${process.env.NODE_ENV},
  );
});
`
	path := projectName + "/src/server.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateAppFile(projectName string) {
	code := `import express from 'express';
import 'express-async-errors';
import bodyParser from 'body-parser';
import cors from 'cors';
import initRoutes from './init/init.routes';
import errorHandler from './middlewares/errorHandler';
const morgan = require('morgan');
const app = express();
app.use(bodyParser.json());

app.use(morgan);
app.get('/ping', (_req, res) => {
  res.send('pong');
});
initRoutes(app);
app.use(errorHandler);
export default app;
`
	path := projectName + "/src/app.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateHttpExceptionFile(projectName string) {
	code := `class HttpException extends Error {
  status: number;

  message: string;

  constructor(status: number, message: string) {
    super(message);
    this.status = status;
    this.message = message;
  }
}

export default HttpException;
`
	path := projectName + "/src/exceptions/httpException.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateInitRoutesFile(projectName string) {
	code := `import { Express } from 'express';
import Controller  from '../shared/interfaces/controller.interface';

const URL_PREFIX = '/api/v1';
export default (app: Express): void => {
  const controllers: Controller[] = [
  ];

  controllers.forEach(controller => {
    app.use(URL_PREFIX + controller.path, controller.route);
  });
};
`
	path := projectName + "/src/init/init.routes.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateDataValidatorFile(projectName string) {
	code := `import { RequestHandler } from 'express';
import { validate, ValidationError } from 'class-validator';
import HttpException from '../exceptions/httpException';

function validationMiddleware<T>(
  ValidationType: new (...args: any[]) => T,
): RequestHandler {
  return async (req, res, next) => {
    const dtoObject = new ValidationType(req.body) as any;
    const errors: ValidationError[] = await validate(dtoObject);

    if (errors.length > 0) {
      const message = errors
        .map((error: ValidationError) => {
          return Object.values(error.constraints || {});
        })
        .join(', ');
      next(new HttpException(400, message));
    } else {
      next();
    }
  };
}

export default validationMiddleware;

`
	path := projectName + "/src/middlewares/dataValidator.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateErrorHanlderFile(projectName string) {
	code := `/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable no-unused-vars */
import { Request, Response, NextFunction } from 'express';
import logger from '../shared/logger';
import HttpException from '../exceptions/httpException';

export default (
  err: HttpException,
  req: Request,
  res: Response,
  next: NextFunction,
): void => {
  logger.error(err);
  console.error({
    Source: req.originalUrl || 'Not provided',
    Message: err.message || 'Not provided',
    stack: err.stack || 'Not provided',
    status: err.status || 'Not provided',
  });
  res.status(err.status || 500).send(err.message);
};
`
	path := projectName + "/src/middlewares/errorHandler.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateAuthFile(projectName string) {
	code := `import jwt from 'jsonwebtoken';
import { Request, Response, NextFunction } from 'express';
import HttpException from '../exceptions/httpException';
import { config } from '../config';

require('dotenv').config();

const secretKey: string = config.JWT_TOKEN_SECRET;

const verifyTokenMiddleware = (
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

export default verifyTokenMiddleware;
`
	path := projectName + "/src/middlewares/auth.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateControllerInterfaceFile(projectName string) {
	code := `import { Router } from 'express';

export default interface Controller {
  path: string;
  route: Router;
  initializeRoutes(): void;
}
`
	path := projectName + "/src/shared/interfaces/controller.interface.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
func CreateLoggerFile(projectName string) {
	code := `import { createLogger, transports, format } from 'winston';

const logger = createLogger({
  format: format.combine(
    format.timestamp({
      format: 'YYYY-MM-DD HH:mm:ss',
    }),
    format.prettyPrint(),
  ),
  transports: [
    new transports.File({
      filename: 'logs/errors.log',
      level: 'error',
    }),
  ],
  exceptionHandlers: [
    new transports.File({
      filename: 'logs/exceptions.log',
    }),
  ],
});

logger.exceptions.handle(
  new transports.Console({
    format: format.combine(format.prettyPrint()),
  }),
);

logger.exitOnError = false;

export default logger;
`
	path := projectName + "/src/shared/logger.ts"
	Helpers.GenerateJavascriptFile(path, code)
}
