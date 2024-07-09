import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate } from 'k6/metrics';
// import { uuidv4 } from "https://jslib.k6.io/k6-utils/1.0.0/index.js";

export let options = {
  stages: [
    { duration: '1m', target: 20 },  // Ramp-up to 20 users
    { duration: '3m', target: 20 },   // Stay at 20 users
    { duration: '1m', target: 0 },   // Ramp-down to 0 users
  ],
  thresholds: {
    http_req_duration: ['p(95)<200'], // 95% requests should be below 200ms
    'http_req_failed{status:200}': ['rate<0.01'], // Less than 1% requests should fail
  },
};

export default function () {
  // Testing API GET list
  let res = http.get('http://localhost:8080/books');
  check(res, {
    'status was 200': (r) => r.status == 200,
  });

  // Testing API GET detail
  res = http.get('http://localhost:8080/books/1');
  check(res, {
    'status was 200': (r) => r.status == 200,
  });

  sleep(1);
}
