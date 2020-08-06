/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */
package integration

import (
	"github.com/wso2/product-apim-tooling/import-export-cli/integration/base"
	"github.com/wso2/product-apim-tooling/import-export-cli/integration/testutils"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
	"path/filepath"
	"testing"
	"time"
)

//Initialize a project Initialize an API without any flag
func TestInitializeProject(t *testing.T) {
	username := superAdminUser
	password := superAdminPassword
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)

	args := &testutils.InitTestArgs{
		CtlUser:  testutils.Credentials{Username: username, Password: password},
		SrcAPIM:  apim,
		InitFlag: projectName,
	}

	testutils.ValidateInitializeProject(t, args)
}

//Initialize an API with --definition flag
func TestInitializeAPIWithDefinitionFlag(t *testing.T) {
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestApiDefinitionPath,
		ForceFlag: false,
	}

	testutils.ValidateInitializeProjectWithDefinitionFlag(t, args)
}

//Initialize an API from Swagger 2 Specification
func TestInitializeAPIFromSwagger2Definition(t *testing.T) {
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestSwagger2DefinitionPath,
		ForceFlag: false,
	}

	testutils.ValidateInitializeProjectWithOASFlag(t, args)
}

//Initialize an API from OpenAPI 3 Specification
func TestInitializeAPIFromOpenAPI3Definition(t *testing.T) {
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPI3DefinitionPath,
		ForceFlag: false,
	}

	testutils.ValidateInitializeProjectWithOASFlag(t, args)
}

//Initialize an API from API Specification URL
func TestInitializeAPIFromAPIDefinitionURL(t *testing.T) {
	username := superAdminUser
	password := superAdminPassword
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPISpecificationURL,
		ForceFlag: false,
	}

	testutils.ValidateInitializeProjectWithOASFlag(t, args)
}

//Import API from initialized project with swagger 2 definition
func TestImportProjectCreatedFromSwagger2Definition(t *testing.T) {
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestSwagger2DefinitionPath,
		ForceFlag: false,
	}

	//Assert that project import to publisher portal is successful
	testutils.ValidateImportInitializedProject(t, args)
}

//Import API from initialized project with openAPI 3 definition
func TestImportProjectCreatedFromOpenAPI3Definition(t *testing.T) {
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPI3DefinitionPath,
		ForceFlag: false,
	}

	//Assert that project import to publisher portal is successful
	testutils.ValidateImportInitializedProject(t, args)
}

//Import API from initialized project from API definition which is already in publisher without --update flag
func TestImportProjectCreatedFailWhenAPIIsExisted(t *testing.T) {
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPI3DefinitionPath,
		ForceFlag: false,
	}

	//Import API for the First time
	testutils.ValidateImportInitializedProject(t, args)

	//Import API for the second time
	testutils.ValidateImportFailedWithInitializedProject(t, args)
}

//Import API from initialized project from API definition which is already in publisher with --update flag
func TestImportProjectCreatedPassWhenAPIIsExisted(t *testing.T) {
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPI3DefinitionPath,
		ForceFlag: false,
	}

	//Import API for the First time
	testutils.ValidateImportInitializedProject(t, args)

	//Import API for the second time
	testutils.ValidateImportUpdatePassedWithInitializedProject(t, args)
}

//Import Api with a Document and Export that Api with a Document
func TestImportAndExportAPIWithDocument(t *testing.T) {
	username := superAdminUser
	password := superAdminPassword
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPI3DefinitionPath,
		ForceFlag: false,
	}

	testutils.ValidateInitializeProjectWithOASFlag(t, args)

	projectPath, _ := filepath.Abs(projectName)
	//Move doc file to created project
	srcPathForDoc, _ := filepath.Abs(utils.TestCase1DocPath)
	destPathForDoc := projectPath + utils.TestCase1DestPathSuffix
	base.Copy(srcPathForDoc, destPathForDoc)

	//Move docMetaData file to created project
	srcPathForDocMetadata, _ := filepath.Abs(utils.TestCase1DocMetaDataPath)
	destPathForDocMetaData := projectPath + utils.TestCase1DestMetaDataPathSuffix
	base.Copy(srcPathForDocMetadata, destPathForDocMetaData)

	//Import the project with Document
	testutils.ValidateImportUpdatePassedWithInitializedProject(t, args)

	testutils.ValidateAPIWithDocIsExported(t, args, utils.DevFirstDefaultAPIName, utils.DevFirstDefaultAPIVersion)
}

//Import Api with an Image and Export that Api with an image (.png Type)
func TestImportAndExportAPIWithPngIcon(t *testing.T) {
	username := superAdminUser
	password := superAdminPassword
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPI3DefinitionPath,
		ForceFlag: false,
	}

	testutils.ValidateInitializeProjectWithOASFlag(t, args)

	//Move icon file to created project
	projectPath, _ := filepath.Abs(projectName)
	srcPathForIcon, _ := filepath.Abs(utils.TestCase2PngPath)
	destPathForIcon := projectPath + utils.TestCase2DestPngPathSuffix
	base.Copy(srcPathForIcon, destPathForIcon)

	//Import the project with icon image(.png)
	testutils.ValidateImportUpdatePassedWithInitializedProject(t, args)

	testutils.ValidateAPIWithIconIsExported(t, args, utils.DevFirstDefaultAPIName, utils.DevFirstDefaultAPIVersion)
}

//Import Api with an Image and Export that Api with an image (.jpeg Type)
func TestImportAndExportAPIWithJpegImage(t *testing.T) {
	apim := apimClients[0]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPI3DefinitionPath,
		ForceFlag: false,
	}

	testutils.ValidateInitializeProjectWithOASFlag(t, args)

	//Move Image file to created project
	projectPath, _ := filepath.Abs(projectName)
	srcPathForImage, _ := filepath.Abs(utils.TestCase2JpegPath)
	destPathForImage := projectPath + utils.TestCase2DestJpegPathSuffix
	base.Copy(srcPathForImage, destPathForImage)

	//Import the project with icon image(.jpeg) provided
	testutils.ValidateImportUpdatePassedWithInitializedProject(t, args)

	testutils.ValidateAPIWithImageIsExported(t, args, utils.DevFirstDefaultAPIName, utils.DevFirstDefaultAPIVersion)
}

//Import and export API with updated thumbnail and document and assert that
func TestUpdateDocAndImageOfAPIOfExistingAPI(t *testing.T) {
	apim := apimClients[1]
	projectName := base.GenerateRandomName(16)
	username := superAdminUser
	password := superAdminPassword

	args := &testutils.InitTestArgs{
		CtlUser:   testutils.Credentials{Username: username, Password: password},
		SrcAPIM:   apim,
		InitFlag:  projectName,
		OasFlag:   utils.TestOpenAPI3DefinitionPath,
		ForceFlag: true,
	}
	//Initialize the project
	testutils.ValidateInitializeProjectWithOASFlag(t, args)

	//Move doc file to created project
	projectPath, _ := filepath.Abs(projectName)
	srcPathForDoc, _ := filepath.Abs(utils.TestCase2DocPath)
	destPathForDoc := projectPath + utils.TestCase2DestPathSuffix
	base.Copy(srcPathForDoc, destPathForDoc)

	//Move Image file to created project
	srcPathForImage, _ := filepath.Abs(utils.TestCase2JpegPath)
	destPathForImage := projectPath + utils.TestCase2DestJpegPathSuffix
	base.Copy(srcPathForImage, destPathForImage)

	//Import the project with Document and image thumbnail
	testutils.ValidateImportUpdatePassedWithInitializedProject(t, args)

	//Update doc file to created project
	srcPathForDocUpdate, _ := filepath.Abs(utils.TestCase1DocPath)
	destPathForDocUpdate := projectPath + utils.TestCase1DestPathSuffix
	base.Copy(srcPathForDocUpdate, destPathForDocUpdate)

	//Update docMetaData file to created project
	srcPathForDocMetadataUpdate, _ := filepath.Abs(utils.TestCase1DocMetaDataPath)
	destPathForDocMetaDataUpdate := projectPath + utils.TestCase1DestMetaDataPathSuffix
	base.Copy(srcPathForDocMetadataUpdate, destPathForDocMetaDataUpdate)

	//Update icon file to created project
	srcPathForIcon, _ := filepath.Abs(utils.TestCase2PngPath)
	destPathForIcon := projectPath + utils.TestCase2DestPngPathSuffix
	base.Copy(srcPathForIcon, destPathForIcon)

	time.Sleep(1 * time.Second)
	//Import the project with updated Document and updated image thumbnail
	testutils.ValidateImportUpdatePassedWithInitializedProject(t, args)

	//Validate that image has been updated
	testutils.ValidateAPIWithDocIsExported(t, args, utils.DevFirstDefaultAPIName, utils.DevFirstDefaultAPIVersion)

	//Validate that document has been updated
	testutils.ValidateAPIWithIconIsExported(t, args, utils.DevFirstDefaultAPIName, utils.DevFirstDefaultAPIVersion)
}
