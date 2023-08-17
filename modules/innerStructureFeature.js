import * as process from 'process';
import * as fs from 'fs';
import * as path from 'path';

import { createFeatureDtoFileContent,
        updateFeatureDtoFileContent,
        featureControllerFileContent,
        featureServiceFileContent,
        featureInterfaceFileContent,
        featureModelFileContent,
        indexFileContent
        } from "../functions/innerStructureFunctions.js";


export async function innerStructureFeature( featureName ){
    const projectRoot = path.join(
        process.cwd()
    );
    // attaque the src/features folder :
    let featuresFolder = path.join(
        projectRoot,'Node_Tuto','src','features'
    );
    console.log(featuresFolder);
    // create the folder feature : 
    const currentFeature  = path.join(
        featuresFolder,featureName
    ).toString();
    fs.mkdirSync(currentFeature);
    // create the dtos folder for the feature : 
    const currentDto = path.join(
        currentFeature,
        "dtos"
    ).toString();
    fs.mkdirSync(
        currentDto
    );
    // create the Update , Create Dto files :
    fs.writeFileSync(
        path.join(
            currentDto,
            `create${featureName}.dto.ts`
        ),
        createFeatureDtoFileContent(featureName)
    );
    fs.writeFileSync(
        path.join(
            currentDto,
            `update${featureName}.dto.ts`
        ),
        updateFeatureDtoFileContent(featureName)
    );
    // create the interface : 
    fs.writeFileSync(
        path.join(
            currentFeature,
            `${featureName}.interface.ts`
        ),
        featureInterfaceFileContent(featureName)
    );
    // create the model : 
    fs.writeFileSync(
        path.join(
            currentFeature,
            `${featureName}.model.ts`
        ),
        featureModelFileContent(featureName)
    )
    // create controller file : 
    fs.writeFileSync(
        path.join(
            currentFeature,
            `${featureName}.controller.ts`
        ),
        featureControllerFileContent(featureName)
    );
    // create the service file :
    fs.writeFileSync(
        path.join(
            currentFeature,
            `${featureName}.services.ts`
        ),
        featureServiceFileContent(featureName)
    )
    // create the index file to export  :
    fs.writeFileSync(
        path.join(
            currentFeature,
            "index.ts"
        ),
        indexFileContent(featureName)
    ); 
}