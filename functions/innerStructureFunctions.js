/**
 * 
 * @returns will create dev.json, staging.json and production.json file
 * @warning changing this function content lead to changing package json file (scripts)
 */
export function envJsonFileContent(){
return {"server": {
    "port": "8084"
  }} 

}
/**
 * 
 * @returns conncet to mongo database function
 */
export function databaseFileContent(){
 return `import config from 'config';
import { MongoClient } from 'mongodb';

const connectDB = async (): Promise<any> => {
  let mongoClient: any = null;

  const databaseName = '';// fill it with you database name

  const mongoUrl: string = config.get(''); // get the url either env json files or .env file 

  const options: Record<string, any> = {
    useNewUrlParser: true,
    useUnifiedTopology: true,
    serverSelectionTimeoutMS: 600000, // Wait for 1 min
  };
  if (mongoClient) {
    return mongoClient.db(databaseName);
  }

  try {
    mongoClient = await MongoClient.connect(mongoUrl, options);
    return mongoClient.db(databaseName);
  } catch (error) {
    return Promise.reject(error);
  }
};

export default connectDB;
`
}
/**
 * 
 * @returns function can be used to return errors to client
 */
export function httpExceptionFileContent(){
    return `class HttpException extends Error {
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
}
/**
 * 
 * @returns file to initalize mount your apis to the app
 */
export function initFileContent(){
    return `import { Express } from 'express';
import { Controller } from '../shared';


const URL_PREFIX_V1 = '/api/v1';
const URL_PREFIX = URL_PREFIX_V1;

export default (app: Express): void => {
  const controllers: Controller[] = [
 // add your controller here 
  ];
// addd midlewares for controllers
  controllers.forEach(controller => {
    app.use(
    
    );
  });
};
`
}
/**
 * 
 * @returns data validator file responsible for validation control on request body
 */
export function dataValidatorFileContent(){
    return `
    import { plainToClass } from 'class-transformer';
import { validate, ValidationError } from 'class-validator';
import * as express from 'express';
import HttpException from '../exceptions/httpException';
import { HttpStatusEnum } from '../shared';

function validationMiddleware<T>(type: any): express.RequestHandler {
  return async (req, res, next): Promise<void> => {
    const errors: ValidationError[] = await validate(
      plainToClass(type, req.body),
    );
    if (errors.length > 0) {
      const message = errors
        .map(
          (error: ValidationError) =>
            error.constraints && Object.values(error.constraints),
        )
        .join(', ');
      next(
        new HttpException(
          HttpStatusEnum.BAD_REQUEST,
          JSON.stringify({ message, requestBody: req.body }),
        ),
      );
    } else {
      next();
    }
  };
}

export default validationMiddleware;
`
}
export function errorHanlderFileContent(){
    return `
import { Request, Response, NextFunction } from 'express';
import logger from '../utils/logger';
import HttpException from '../exceptions/httpException';
import { HttpStatusEnum } from '../shared';

export default (
  err: HttpException,
  req: Request,
  res: Response,
  next: NextFunction,
): void => {
  logger.error(err);
  console.error('Error Occured ', {
    Source: req.originalUrl,
    Message: err.message || 'Not provided',
    stack: err.stack || 'Not provided',
    status: err.status || 'Not provided',
  });
  res.status(err.status || HttpStatusEnum.SERVER_ERROR).send(err.message);
};
`
}
/**
 * 
 * @returns enum for http status code 
 */
export function enumFileContent(){
    return `export enum HttpStatusEnum {
  SUCCESS = 200,
  CREATED = 201,
  SUCCESS_NO_CONTENT = 204,
  BAD_REQUEST = 400,
  UNAUTHORIZED = 401,
  FORBIDDEN = 403,
  NOT_FOUND = 404,
  SERVER_ERROR = 500,
}
`
}
/**
 * 
 * @returns interface to fetch mongodb vars connection
 */
export function configInterfaceFileContent(){
    return `export interface DBConfig {
  port: string;
  host: string;
  name: string;
  user: string;
  password: string;
}
`
}
/**
 * @returns controller interface 
 */
export function controllerInterfaceFileContent(){
  return `import { Router } from 'express';

export default interface Controller {
  path: string;
  route: Router;
  initializeRoutes(): void;
}
`
}
/**
 * 
 * @description a file to export the shared interface to be used anywhere in the project
 */
export function indexForInterfacesFileContent(){
    return `export * from './Enums/httpStatus.enum';
export { default as Controller } from './interfaces/controller.interface';
`
}
/**
 * @description file to specify custom function signature
 */
export function typesFileContent(){
return `export declare type LangType = {
  en?: string | undefined;
  fr?: string | undefined;
  ar?: string | undefined;
};

export declare type ResponseStatusType = {
  status: boolean;
  message: LangType;
};

export declare type PaginationType = {
  take?: string | undefined;
  skip?: string | undefined;
};

export declare type ListResult<Type> = {
  list: Type[];
  count: number;
};
`
}
/**
 * @description logger file 
 */
export function loggerFileContent(){
    return `import { createLogger, transports, format } from 'winston';

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
}
export function appFileContent(){
   return `
import express, { Request, Response } from 'express';
import 'express-async-errors';
import bodyParser from 'body-parser';
import cors from 'cors';
import initRoutes from './init/init.routes';
import errorHandler from './middlewares/errorHandler';
import morgan from 'morgan';

const app = express();
app.use(bodyParser.json());
app.use(cors({ origin: '*' }));

app.use(
  morgan(function(tokens:any, req, res) {
    // Get the request body as a string
    const requestBody = JSON.stringify(req.body);
    // Define log colors for different log levels
    const statusColor =
      tokens.status(req, res) >= 500 ? '\\x1b[31m' : '\\x1b[32m'; // Red for 5xx, Green for 2xx
    const resetColor = '\\x1b[0m'; // Reset color to default

    // Check if the request path matches the desired paths
    const isDesiredRequest =
      req.url.startsWith('/paidTrip') || req.url.startsWith('/coupons');

    // If the request path is not one of the desired paths, skip logging
    if (!isDesiredRequest) {
      return;
    }

    // Construct the log message
    const logMessage = [
      '==== new request ====',
      \`\${tokens.method(req, res)} \${tokens.url(req, res)}\`,
      \`\${statusColor}\${tokens.status(req, res)}\${resetColor}\`,
      \`Body: \${requestBody}\`,
    ].join(' ');

    return logMessage;
  })
);

app.get('/ping', (_req, res) => {
  res.send('pong');
});

initRoutes(app);

app.use(errorHandler);

export default app;
`;
}
export function serverFileContent(){
    return `import config from 'config';
import app from './app';
require('dotenv').config();
const serverConfig: { port: number } = config.get('server');
app.listen(serverConfig.port, () => {
  console.info(\`Listening on port \${serverConfig.port}\`);
});
`
}


export function createFeatureDtoFileContent(featureName){
  return `
  import { IsNumber,
    IsArray,
    IsNotEmpty,
    IsString,
    IsDate,
    IsBoolean,
    IsEnum
} from "class-validator";

export default class Create${featureName}Dto {
};
  `;
}

export function updateFeatureDtoFileContent(featureName){
  return `
  import { IsNumber,
    IsArray,
    IsString,
    IsDate,
    IsBoolean,
    IsEnum
} from "class-validator";

export default class Update${featureName}Dto {
};
  `;
}



export function featureControllerFileContent(featureName){
  return `
  import { Router, Response, Request } from 'express';
  import { Controller } from '../../shared';
  import ${featureName}Service from './${featureName}.services';
  import validationMiddleware from '../../middlewares/dataValidator';
  import { authMiddleware } from '../../middlewares/auth';
  import Create${featureName}Dto from './dtos/create${featureName}.dto';
  import Update${featureName}Dto from './dtos/update${featureName}.dto';
  import { CustomRequest } from '../../shared/interfaces/customRequest.interface';

  export default class ${featureName}Controller implements Controller {
  path="/${featureName.toLowerCase()}";
  route = Router();
  constructor(){
      this.initializeRoutes();
  };
  initializeRoutes():void{}
}
  ` 
}


export function featureServiceFileContent(featureName){
  return `
  import HttpException from "../../exceptions/httpException";
  import ${featureName}Model from "./${featureName}.model";
  import I${featureName} from "./${featureName}.interface";
  import Update${featureName}Dto from "./dtos/update${featureName}.dto";
  import Create${featureName}Dto from "./dtos/create${featureName}.dto";

  export default class ${featureName}Service{ 
    // create : 
    static create${featureName} = async (${featureName}Payload:Create${featureName}Dto):Promise<I${featureName}> => {
        try{
            // code ......
        }catch( error ){
            console.log( error );
            throw new HttpException(500, "couldnt create ${featureName}");
        }
    };
    // update : 
    static update${featureName} = async (${featureName}Id:string , ${featureName}Payload:Update${featureName}Dto):Promise<I${featureName}> => {
        try{
            // code ... 
        }catch( error ){
            console.log( error );
            throw new HttpException(500, "couldnt update ${featureName}");
        }       
    };
    // get all the events : 
    static get${featureName}List = async ():Promise<I${featureName}[]> =>{
        const ${featureName}Results:I${featureName}[] = [];
        const ${featureName}List = await ${featureName}Model.find();
        ${featureName}List.forEach(
          (item: I${featureName}) => {
              ${featureName}Results.push(item);
            }
        )
        return ${featureName}Results;
    }
}; `;

} 

export function featureInterfaceFileContent(featureName){
  return `
  import mongoose , { Document } from "mongoose";
  export default interface I${featureName} extends Document {
        // attributes ..
  };
  `;
}


export function featureModelFileContent(featureName){
  return `
import mongoose from "mongoose";
import I${featureName} from "./event.interface";

const ${featureName}Schema = new mongoose.Schema<IEvent>();


${featureName}Schema.set('timestamps',true).toString();
// create a mongodb model based on the schema : 
const ${featureName}Model = mongoose.model<I${featureName}>("${featureName}",${featureName}Schema);
export default ${featureName}Model;
  `;
}

export function indexFileContent(featureName){
  return `
  export { default as ${featureName}Controller } from './${featureName}.controller';
  export { default as ${featureName}Service } from './${featureName}.services';
  export { default as ${featureName} } from './${featureName}.model';
  export { default as I${featureName} } from './${featureName}.interface';
  `;
}