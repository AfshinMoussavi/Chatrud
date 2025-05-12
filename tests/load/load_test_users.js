import http from 'k6/http';
import { check } from 'k6';

export const options = {
    scenarios: {
        constant_request_rate: {
            executor: 'constant-arrival-rate',
            rate: 1000,
            timeUnit: '1s',
            duration: '1m',
            preAllocatedVUs: 400,
            maxVUs: 1000,
        },
    },
};

export default function () {
    const res = http.get('http://localhost:8080/api/auth/users');
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response contains users': (r) => Array.isArray(r.json().data),
    });
}
