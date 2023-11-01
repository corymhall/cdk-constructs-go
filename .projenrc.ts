import { awscdk } from 'projen';

const props: Omit<awscdk.AwsCdkConstructLibraryOptions, 'name'> = {
  author: 'corymhall',
  authorAddress: '43035978+corymhall@users.noreply.github.com',
  cdkVersion: '2.103.1',
  defaultReleaseBranch: 'main',
  jsiiVersion: '~5.0.0',
  projenrcTs: true,
  repositoryUrl: 'https://github.com/corymhall/cdk-constructs-go.git',
  githubOptions: {
    mergify: false,
  },
  releaseToNpm: false,
  devDeps: [
    'cdk-import',
  ],
};
const project = new awscdk.AwsCdkConstructLibrary({
  ...props,
  name: 'cdk-constructs-go',
});

new awscdk.AwsCdkConstructLibrary({
  outdir: 'datadog',
  parent: project,
  name: 'cdk-constructs-datadog-go',
  publishToGo: {
    moduleName: 'github.com/corymhall/cdk-constructs-go',
    packageName: 'cdkdatadog',
    gitBranch: 'constructs',
  },
  ...props,
});

project.postCompileTask.spawn(project.addTask('datadog-build', {
  cwd: 'datadog',
  exec: 'rm -rf test && yarn projen build',
}));

project.release?.publisher.publishToGo({
  prePublishSteps: [
    {
      name: 'Prepare Repository',
      run: 'mv dist .repo',
    },
    {
      name: 'Install Dependencies',
      run: 'cd .repo && yarn install --check-files --frozen-lockfile',
    },
    {
      name: 'Create go artifact',
      run: 'cd .repo/datadog && npx projen package:go',
    },
    {
      name: 'Collect go Artifact',
      run: 'mv .repo/datadog/dist dist',
    },
  ],
  githubTokenSecret: 'GITHUB_TOKEN',
  gitBranch: 'constructs',
});

project.gitignore.exclude('/datadog/test/hello.test.ts');

project.synth();
