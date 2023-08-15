import * as process from 'process';
import * as fs from 'fs';
import * as path from 'path';
import {
  envJsonFileContent,
  databaseFileContent,
  httpExceptionFileContent,
  initFileContent,
  dataValidatorFileContent,
  errorHanlderFileContent,
  enumFileContent,
  configInterfaceFileContent,
  controllerInterfaceFileContent,
  indexForInterfacesFileContent,
  typesFileContent,
  loggerFileContent,
  appFileContent,
  serverFileContent
} from "../functions/innerStructureFunctions.js";


export async function innerStructureModule(projectName){
const projectRoot = path.join(process.cwd(), projectName.toString());
// Create the src folder
    fs.mkdirSync(path.join(projectRoot, 'src'));
// Create config folder 
    fs.mkdirSync(path.join(projectRoot, 'src','config'));
// Create env files inside config folder
const envJsonFile = envJsonFileContent();
    fs.writeFileSync(path.join(projectRoot,'src','config','dev.json'),JSON.stringify(envJsonFile,null,2));
    fs.writeFileSync(path.join(projectRoot,'src','config','staging.json'),JSON.stringify(envJsonFile,null,2));
    fs.writeFileSync(path.join(projectRoot,'src','config','production.json'),JSON.stringify(envJsonFile,null,2));
// Create db folder inside config folder
    fs.mkdirSync(path.join(projectRoot, 'src','config','db'));
// Create database folder and the database file inside config folder
const databaseFile = databaseFileContent()
    fs.writeFileSync(path.join(projectRoot,'src','config',"db","database.ts"),databaseFile)
// Crreate exceptions folder
    fs.mkdirSync(path.join(projectRoot, 'src','exceptions'));
// Create file inside exceptions folder
const httpExceptionFile = httpExceptionFileContent()
    fs.writeFileSync(path.join(projectRoot,'src','exceptions',"httpException.ts"),httpExceptionFile)
// Create init folder
    fs.mkdirSync(path.join(projectRoot, 'src','init'));
// Create features folder 
    fs.mkdirSync(path.join(projectRoot, 'src','features'));
// Create init file 
const initFile = initFileContent()
    fs.writeFileSync(path.join(projectRoot,'src','init',"init.routes.ts"),initFile)
    // Create middlewares folder
    fs.mkdirSync(path.join(projectRoot, 'src','middlewares'));
// Create file inside middleware folder
const dataValidatorFile = dataValidatorFileContent();
const errorHanlderfile = errorHanlderFileContent();
    fs.writeFileSync(path.join(projectRoot,'src','middlewares',"dataValidator.ts"),dataValidatorFile)
    fs.writeFileSync(path.join(projectRoot,'src','middlewares',"errorHandler.ts"),errorHanlderfile)
// Create shared folder
    fs.mkdirSync(path.join(projectRoot, 'src','shared'));
// Create Dtos folder inside shared folder (used for middleware dtos usually)
    fs.mkdirSync(path.join(projectRoot, 'src','shared','Dtos'));
// Create Enums folder inside shared folder
    fs.mkdirSync(path.join(projectRoot, 'src','shared','Enums'));
// Create file inside enums folder
const httpStatusEnumFile = enumFileContent();
    fs.writeFileSync(path.join(projectRoot,'src','shared',"Enums","httpStatus.enum.ts"),httpStatusEnumFile)
// Create interfaces folder inside shared folder
    fs.mkdirSync(path.join(projectRoot, 'src','shared','interfaces'));
// Create config.interface file and controller interface inside interfaces folder
const configInterfaceFile = configInterfaceFileContent() 
const controllerInterfaceFile = controllerInterfaceFileContent()
    fs.writeFileSync(path.join(projectRoot, 'src','shared','interfaces','config.interface.ts'),configInterfaceFile);
    fs.writeFileSync(path.join(projectRoot, 'src','shared','interfaces','controller.interface.ts'),controllerInterfaceFile);
// Create index and types files inside shared folder 
const indexForInterfacesFile = indexForInterfacesFileContent();
const typesFile = typesFileContent()
    fs.writeFileSync(path.join(projectRoot, 'src','shared','index.ts'),indexForInterfacesFile);
    fs.writeFileSync(path.join(projectRoot, 'src','shared','types.ts'),typesFile)
// create test folder ( it's empty folder, in the future i will add the files maybe)
    fs.mkdirSync(path.join(projectRoot, 'src','tests'));
// Create utils folder
    fs.mkdirSync(path.join(projectRoot, 'src','utils'));
// Create logger file inside the utils folder
const loggerFile = loggerFileContent()
    fs.writeFileSync(path.join(projectRoot, 'src','utils','logger.ts'),loggerFile);
// Create app file inside src folder 
const appfile = appFileContent();
    fs.writeFileSync(path.join(projectRoot, 'src','app.ts'),appfile)
    // Create server file inside src folder 
const serverFile = serverFileContent()
    fs.writeFileSync(path.join(projectRoot, 'src','server.ts'),serverFile);

}