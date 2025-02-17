import counter from 'k6/x/counter'
import harbor from 'k6/x/harbor'
import { Rate } from 'k6/metrics'

const missing = Object()

function getEnv(env, def = missing) {
    const value = __ENV[env] ? __ENV[env] : def
    if (value === missing) {
        throw (`${env} envirument is required`)
    }

    return value
}

const teardownResources = getEnv('TEARDOWN_RESOURCES', 'true') === 'true'

export let errorRate = new Rate('errors');

export let options = {
    setupTimeout: '1h',
    teardownTimeout: '30m',
    noUsageReport: true,
    vus: 500,
    iterations: 1000,
    thresholds: {
        'iteration_duration{scenario:default}': [
            `max>=0`,
        ],
        'iteration_duration{group:::setup}': [`max>=0`],
        'iteration_duration{group:::teardown}': [`max>=0`]
    }
};

export function setup() {
    harbor.initialize({
        scheme: getEnv('HARBOR_SCHEME', 'https'),
        host: getEnv('HARBOR_HOST'),
        username: getEnv('HARBOR_USERNAME', 'admin'),
        password: getEnv('HARBOR_PASSWORD', 'Harbor12345'),
        insecure: true,
    })

    return {
        now: Date.now(),
    }
}

export default function ({ now }) {
    const i = counter.up() - 1

    const projectName = `project-${now}-${i}`

    try {
        harbor.createProject({ projectName })
    } catch (e) {
        console.log(e)
        errorRate.add(true)
    }
}

export function teardown({ now }) {
    if (teardownResources) {
        const pageSize = 15

        while(true) {
            try {
                const { projects } = harbor.listProjects({name: `project-${now}-`, pageSize})

                for (const project of projects) {
                    try {
                        harbor.deleteProject(project.name)
                    } catch (e) {
                        console.log(e)
                    }
                }

                if (projects.length < pageSize) {
                    break
                }
            } catch (e) {
                console.log(e)
                break
            }
        }
    }
}
