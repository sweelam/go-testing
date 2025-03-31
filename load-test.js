import http from 'k6/http'
import { check, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '1m', target: 1000 },   // Ramp-up to 100 VUs
    { duration: '1m', target: 1000 },   // Stay at 100 VUs for 1 minute
    { duration: '30s', target: 0 },     // Ramp-down to 0 VUs
  ],
};

export default function () {
  let res = http.get('http://localhost:8080/api/v1/jobs');
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(1);
}
