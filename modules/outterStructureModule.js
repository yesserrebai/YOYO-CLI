import * as process from 'process';
import * as fs from 'fs';
import * as path from 'path';
import{
 nodeMonFileContent,packageJsonFileContent,tsConfigFileContent,jestConfigFileContent,dockerFileContent,copyEnvFileContent,prettierFileContent,gitIgnoreFileContent,eslintrcFileContent,eslintIgnoreFileContent
} from '../functions/outterStrcutureFunctions.js'

export async function outterStructureModule(projectName){
    const projectRoot = path.join(process.cwd(), projectName.toString());
    // Create the project folder if it doesn't exist
if (!fs.existsSync(projectRoot)) {
    fs.mkdirSync(projectRoot);
}
    // Create package.json file 
    const packageJsonFile = packageJsonFileContent(projectName.toString());
    fs.writeFileSync(path.join(projectRoot,'package.json'),JSON.stringify(packageJsonFile,null,2));
    // Create nodemon file
    const nodeMonFile = nodeMonFileContent()
    fs.writeFileSync(path.join(projectRoot,'nodemon.json'),JSON.stringify(nodeMonFile,null,2));
    // Create tsConfig file 
    const tsConfigFile = tsConfigFileContent();
    fs.writeFileSync(path.join(projectRoot,'tsconfig.json'),JSON.stringify(tsConfigFile,null,2));
    // Create jest config file
    const jestConfigFile = jestConfigFileContent();
    fs.writeFileSync(path.join(projectRoot,'jest.config.json'),JSON.stringify(jestConfigFile,null,2));
    // Create docker file
    const dockerFile = dockerFileContent()
    fs.writeFileSync(path.join(projectRoot,'Dockerfile'),dockerFile);
    // Create copyEnv file 
    const copyEnvFile = copyEnvFileContent();
    fs.writeFileSync(path.join(projectRoot,'copyEnv.js'),copyEnvFile);
    // Create prettier file 
    const prettierFile = prettierFileContent();
    fs.writeFileSync(path.join(projectRoot,'.prettierrc'),prettierFile);
    // Create git ingore file 
    const gitIgnoreFile = gitIgnoreFileContent()
    fs.writeFileSync(path.join(projectRoot,'.gitignore'),gitIgnoreFile);
    // Create eslintrc file
    const eslintrcFile = eslintrcFileContent()
    fs.writeFileSync(path.join(projectRoot,'.eslintrc'),eslintrcFile);
    // Create elsint ignore file 
    const eslintIgnoreFile = eslintIgnoreFileContent();
    fs.writeFileSync(path.join(projectRoot,'.eslintignore'),eslintIgnoreFile);
}