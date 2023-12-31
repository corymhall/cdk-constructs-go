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

// project.gitignore.addPatterns('cdk-monitoring-constructs');

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

project.postCompileTask.reset();

project.postCompileTask.spawn(project.addTask('datadog-build', {
  cwd: 'datadog',
  exec: 'rm -rf test && yarn compile',
}));

project.postCompileTask.spawn(project.addTask('monitoring-build', {
  steps: [
    {
      exec: 'rm -rf cdk-monitoring-constructs && git clone https://github.com/cdklabs/cdk-monitoring-constructs && rm -rf cdk-monitoring-constructs/.git',
    },
    {
      cwd: 'cdk-monitoring-constructs',
      exec: 'yarn install --frozen-lockfile',
    },
    {
      exec: 'jq \'.jsii.targets.go={"moduleName":"github.com/corymhall/cdk-constructs-go","packageName":"cdkmonitoringconstructs"}\' package.json > package.json.tmp && mv package.json.tmp package.json',
      cwd: 'cdk-monitoring-constructs',
    },
    {
      cwd: 'cdk-monitoring-constructs',
      exec: 'yarn compile',
    },
  ],
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
      run: 'ls -la .repo && mv .repo/datadog/dist dist',
    },
    {
      name: 'Create go artifact 2',
      run: 'cd .repo/cdk-monitoring-constructs && yarn install --frozen-lockfile && npx jsii-pacmak --target go --force-target',
    },
    {
      name: 'Collect go Artifact 2',
      run: 'mv .repo/cdk-monitoring-constructs/dist/go/cdkmonitoringconstructs dist/go',
    },
  ],
  gitBranch: 'constructs',
});

project.gitignore.exclude('/datadog/test/hello.test.ts');

project.synth();
