#!/usr/bin/env node

import process from 'process';
import chalk from 'chalk';
import pkg from 'inquirer';
const { createPromptModule  } = pkg;
import { outterStructureModule } from '../modules/outterStructureModule.js';
import {innerStructureModule} from '../modules/innerStructureModule.js';

const { green, yellow, red, blue } = chalk;

const success = green.bold;
const warning = yellow.bold;
const error = red.bold;
const info = blue;

// main is the entry point for set-up cmd
async function main() {
  console.log(info("Welcome to YOYO-CLI!"));

  let prompt = await createPromptModule()([
    {
      type: "input",
      name: "confirmation",
      message: warning("I'm going to create a new project, confirm? (y/n)")
    }
  ]);

  if (prompt.confirmation !== 'y' && prompt.confirmation !== 'yes') {
    console.log(
      info("Aborted project setup.")
      );
    process.exit(0);
  }

  prompt = await createPromptModule()([
    {
      type: "input",
      name: "projectName",
      message: "Enter project name"
    }
  ]);

  console.log(
    success(
      `Creating the outer structure for project: ${prompt.projectName}`
    )
  );
  await outterStructureModule(
    prompt.projectName
  );
  console.log(
    success("Outer structure created successfully.")
  );

  console.log(
    success(`Creating the inner structure for project: ${prompt.projectName}`
    )
  );

  await innerStructureModule(prompt.projectName);

  console.log(
    success("Inner structure created successfully.")
  );

  console.log(
    success('Project setup complete.')
  );
  console.log(
    info('Go to your project and run ') + 
    warning('yarn install') + 
    info(' to install dependencies.')
  );
}

main();
