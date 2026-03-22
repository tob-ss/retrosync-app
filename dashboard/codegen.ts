import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
    schema: './src/lib/assets/schema.graphql',
    documents: './src/lib/assets/genqlient.graphql',
    generates: {
        './graphql/generated.ts': {
            plugins: ['typescript', 'typescript-operations', 'typed-document-node', '@kitql/graphql-codegen']
        }
    }
}

export default config