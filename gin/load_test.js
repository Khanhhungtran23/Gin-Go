import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '30s', target: 10 },  // Ramp-up to 10 users
    { duration: '1m', target: 10 },   // Stay at 10 users
    { duration: '30s', target: 0 },   // Ramp-down to 0 users
  ],
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
