
/**
 * 
 * @returns nodemon.json file
 */
export function nodeMonFileContent(){
    return {
  "watch": ["src"],
  "ext": "ts",
  "exec": "ts-node ./src/server.ts",
  "env": {
    "NODE_ENV": "dev",
    "NODE_CONFIG_DIR": "./src/config"
  }
}
}
/**
 * 
 * @param {string} projectName 
 * @returns the package.json file 
 */
export function packageJsonFileContent(projectName){
   return  {
  "name": projectName,
  "version": "1.0.0",
  "main": "",
  "license": "MIT",
  "dependencies": {
    "axios": "^0.26.1",
    "body-parser": "^1.19.0",
    "class-transformer": "^0.3.1",
    "class-validator": "^0.14.0",
    "config": "^3.2.2",
    "copyfiles": "^2.4.1",
    "cors": "^2.8.5",
    "cross-env": "7.0.3",
    "dotenv": "^16.3.1",
    "express": "^4.17.1",
    "express-async-errors": "^3.1.1",
    "fs-extra": "^11.1.1",
    "jsonwebtoken": "^9.0.0",
    "localtunnel": "^2.0.2",
    "mongodb": "^5.7.0",
    "mongoose": "^7.4.2",
    "morgan": "^1.10.0",
    "winston": "^3.2.1"
  },
  "devDependencies": {
    "@types/config": "^0.0.34",
    "@types/cors": "^2.8.6",
    "@types/express": "^4.17.1",
    "@types/glob": "^8.0.1",
    "@types/jest": "^24.0.18",
    "@types/jsonwebtoken": "^8.3.5",
    "@types/morgan": "^1.9.4",
    "@types/node": "12.7.2",
    "@types/supertest": "^2.0.12",
    "@types/swagger-ui-express": "^4.1.2",
    "@types/validator": "^13.7.2",
    "@typescript-eslint/eslint-plugin": "^2.1.0",
    "@typescript-eslint/parser": "^2.1.0",
    "eslint": "^6.3.0",
    "eslint-config-airbnb-base": "14.0.0",
    "eslint-config-prettier": "^6.2.0",
    "eslint-plugin-import": "^2.18.2",
    "eslint-plugin-prettier": "^3.1.0",
    "jest": "^27.5.1",
    "nodemon": "^2.0.4",
    "prettier": "^1.18.2",
    "supertest": "^6.3.3",
    "ts-jest": "^27.1.4",
    "ts-node": "^8.3.0",
    "typescript": "4.4.2"
  },
  "scripts": {
    "start": "nodemon -L ",
    "start:prod": "export NODE_CONFIG_DIR=./src/config && export NODE_ENV=prod && node dist/server.js",
    "start:stg": "export NODE_CONFIG_DIR=./src/config && export NODE_ENV=staging && node dist/server.js",
    "build": "tsc && npm run copy-config",
    "copy-config": "copyfiles -u 1 \"src/config/*.json\" dist/ && cross-env node ./copyEnv.js",
    "lint": "eslint --fix src/**/*.ts ",
    "test": "export NODE_ENV=dev && export NODE_CONFIG_DIR=./src/config && jest",
    "test:dev": "export NODE_ENV=dev && export NODE_CONFIG_DIR=./src/config && jest --watchAll"
  }
}
}
/**
 * @returns ts config file content e.g configurations
 */
export function tsConfigFileContent(){
 return {
  "include": ["src/**/*"],
  "compilerOptions": {
    "module": "commonjs",
    "target": "es2017",
    "lib": ["es2017", "DOM"],
    "skipLibCheck": true,
    "esModuleInterop": true,
    "experimentalDecorators": true,
    "emitDecoratorMetadata": true,
    "outDir": "dist",
    "strict": true,
    "sourceMap": true,
    "resolveJsonModule": true,
    "strictPropertyInitialization": false
  }
}
} 
/**
 * @returns jest config file content
 */
export function jestConfigFileContent(){
    return `
module.exports = {
  globals: {
    'ts-jest': {
      tsconfig: 'tsconfig.json',
    },
  },
  moduleFileExtensions: ['ts', 'js'],
  transform: {
    '^.+\\.ts$': 'ts-jest',
  },
  testMatch: ['<rootDir>/src/tests/**/*.test.(ts|js)'],
  testEnvironment: 'node',
};
`;
}
/**
 * 
 * @returns dokcer file content
 */
export function dockerFileContent(){
  return `
# This file is for staging. DO NOT use it for production env
FROM node:16-alpine as build

WORKDIR /usr/src/app

ENV NODE_CONFIG_DIR=./src/config

# Install app dependencies
COPY package.json ./

RUN yarn install

# Bundle app source
COPY . .
RUN yarn build -v 

EXPOSE 8084
CMD [ "yarn", "start:stg" ]
`;
}
/**
 * 
 * @returns env file content
 */
export function copyEnvFileContent(){
  return `const fs = require('fs-extra');
          //copy .env to the dist directory
          fs.copySync('.env', 'dist/.env');`
}
/**
 * 
 * @returns prettier file content 
 */
export function prettierFileContent(){
  return `{
  "printWidth": 80,
  "tabWidth": 2,
  "useTabs": false,
  "semi": true,
  "singleQuote": true,
  "trailingComma": "all"
}
`
}
/**
 * 
 * @returns git ignore file content
 */
export function gitIgnoreFileContent(){
  return`node_modules
        dist
        env.prod.json
        routes.ts
        swagger.json
        *.log
        TODO
        .idea
        .vscode
        .DS_Store
        config/default.json
        config/production.json
        dep
        .env`
}
/**
 * 
 * @returns eslint file content
 */
export function eslintrcFileContent(){
  return `{
  "parser": "@typescript-eslint/parser",
  "extends": [
    "plugin:@typescript-eslint/recommended",
    "airbnb-base",
    "prettier"
  ],
  "plugins": ["prettier"],
  "settings": {
    "import/resolver": {
      "node": {
        "extensions": [".js", ".ts"]
      }
    }
  },
  "rules": {
    "no-console": "off",
    "camelcase": "off",
    "@typescript-eslint/camelcase": "off",
    "class-methods-use-this": "off",
    "no-explicit-any": "off",
    "no-throw-literal": "off",
    "no-underscore-dangle": "off",
    "@typescript-eslint/no-explicit-any": "off",
    "import/prefer-default-export": "off",
    "@typescript-eslint/interface-name-prefix": "off",
    "import/extensions": [
      "error",
      "ignorePackages",
      {
        "js": "never",
        "jsx": "never",
        "ts": "never",
        "tsx": "never"
      }
    ]
  },
  "env": {
    "jest": true
  }
}
`
}
/**
 * 
 * @returns eslintIngore file content
 */
export function eslintIgnoreFileContent(){
  return `dist/**/*
          routes.ts`
}