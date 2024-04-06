package crud

const controllerTemplate = `
import Controller from "../../shared/interfaces/controller.interface";
import {{.CapitalizedFeatureName}}Service from "./{{ .FeatureNameLower }}.service";
import { Request, Response, Router } from 'express';

export default class {{ .FeatureName }}Controller implements Controller {
  path = '/{{ .FeatureNameLower }}s';

  route = Router();

  constructor() {
    this.initializeRoutes();
  }

  initializeRoutes(): void {
	this.route.get(
		'',
		this.get{{.CapitalizedFeatureName}}
	);
	this.route.get(
		'/:id',
		this.getOne{{.CapitalizedFeatureName}}
	);
	this.route.post(
      '',
      this.create{{ .CapitalizedFeatureName }},
    );
	this.route.patch(
		'/:id',
		this.update{{.CapitalizedFeatureName}}
	)
	this.route.delete(
		'/:id',
		this.delete{{.CapitalizedFeatureName}}
	)
  }
async get{{ .CapitalizedFeatureName }}(req: Request, res: Response): Promise<void> {
    const body = req.body
      const result = await {{ .CapitalizedFeatureName }}Service.get{{ .CapitalizedFeatureName }}(body);
      res.status(200).send(result)
  }
async getOne{{ .CapitalizedFeatureName }}(req: Request, res: Response): Promise<void> {
    const body = req.body
	const id=req.params.id
      const result = await {{ .CapitalizedFeatureName }}Service.getOne{{ .CapitalizedFeatureName }}(body);
      res.status(200).send(result)
  }

async create{{ .CapitalizedFeatureName }}(req: Request, res: Response): Promise<void> {
    const body = req.body
      const result = await {{ .CapitalizedFeatureName }}Service.create{{ .CapitalizedFeatureName }}(body);
      res.status(200).send(result)
  }
  async update{{ .CapitalizedFeatureName }}(req: Request, res: Response): Promise<void> {
    const body = req.body
	const id=req.params.id
      const result = await {{ .CapitalizedFeatureName }}Service.update{{ .CapitalizedFeatureName }}(body);
      res.status(200).send(result)
  }
  async delete{{ .CapitalizedFeatureName }}(req: Request, res: Response): Promise<void> {
    const body = req.body
	const id=req.params.id
      const result = await {{ .CapitalizedFeatureName }}Service.delete{{ .CapitalizedFeatureName }}(body);
      res.status(200).send(result)
  }
}
`
const serviceTemplate = `
export default class {{ .CapitalizedFeatureName }}Service{
  static create{{ .CapitalizedFeatureName }}=async(body)=>{
    await {{ .CapitalizedFeatureName }}Helper.create{{ .CapitalizedFeatureName }}(body)
    throw 'not implemented'
  }
  static get{{ .CapitalizedFeatureName }}=async(body)=>{
    await {{ .CapitalizedFeatureName }}Helper.get{{ .CapitalizedFeatureName }}(body)
    throw 'not implemented'
  }
  static getOne{{ .CapitalizedFeatureName }}=async(body)=>{
    await {{ .CapitalizedFeatureName }}Helper.getOne{{ .CapitalizedFeatureName }}(body)
    throw 'not implemented'
  }
  static update{{ .CapitalizedFeatureName }}=async(body)=>{
    await {{ .CapitalizedFeatureName }}Helper.update{{ .CapitalizedFeatureName }}(body)
    throw 'not implemented'
  }
  static delete{{ .CapitalizedFeatureName }}=async(body)=>{
    await {{ .CapitalizedFeatureName }}Helper.delete{{ .CapitalizedFeatureName }}(body)
    throw 'not implemented'
  }
}
`
const helperTemplate = `export default class {{ .CapitalizedFeatureName }}Helper{
  static create{{ .CapitalizedFeatureName }}=async(body)=>{
    throw 'not implemented'
  }
  static get{{ .CapitalizedFeatureName }}=async(body)=>{
    throw 'not implemented'
  }
  static getOne{{ .CapitalizedFeatureName }}=async(body)=>{
    throw 'not implemented'
  }
  static update{{ .CapitalizedFeatureName }}=async(body)=>{
    throw 'not implemented'
  }
  static delete{{ .CapitalizedFeatureName }}=async(body)=>{
    throw 'not implemented'
  }
}`
