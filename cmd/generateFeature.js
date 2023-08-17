
import process from 'process';
import chalk from 'chalk';
import pkg from 'inquirer';
const { createPromptModule  } = pkg;
import { innerStructureFeature } from '../modules/innerStructureFeature.js';

const { green, yellow, red, blue } = chalk;

const success = green.bold;
const warning = yellow.bold;
const error = red.bold;
const info = blue;
import { isValidString } from "../utils/validate.js";


const main = async () => {
    if (process.argv.length !== 3) {
        console.log(
            warning("Usage: generate-feature <FeatureName> \n"),
            error("FeatureName is required ")
        );
        process.exit(1); // Exit the script with an error code
    }
    
    // Access the provided parameter
    const featureName  = process.argv[2];
    if ( !isValidString(featureName)){
        console.log(
            error("Error: FeatureName must be a valid string \n")
        );
        process.exit(1);
    }
    // check the feature name : 
    console.log(
        success(
          `Creating the sub-folder for : ${featureName}`
        )
      );
    // generate the new feature folder : 
    await innerStructureFeature( featureName );
    console.log(
        success(
          `${featureName} Created successfully\n`
        ),
        info(
            `Please make sure you add the following lines into:
            
            ./src/init/init.routes.ts
            
            // import the feature controller:
            
            import { ${featureName}Controller } from '../features/${featureName}';
            
            // update the controller array :
            
            const controllers: Controller[] = [
                // push the feature controller: 
                new ${featureName}Controller(),
              ];
            
            // update the api middlewares by adding the following line of code: 
            
            controllers.forEach(controller => {
                app.use(URL_PREFIX + controller.path, controller.route);
              });`
        )
    );
};

main();

