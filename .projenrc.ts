import { awscdk } from 'projen';
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'corymhall',
  authorAddress: '43035978+corymhall@users.noreply.github.com',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'main',
  jsiiVersion: '~5.0.0',
  name: 'cdk-constructs-go',
  projenrcTs: true,
  repositoryUrl: 'https://github.com/corymhall/cdk-constructs-go.git',

  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();