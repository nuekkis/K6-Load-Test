import http from "k6/http";
import { check, sleep } from "k6";
import { SharedArray } from "k6/data";

const config = new SharedArray("config", function() {
    return JSON.parse(open("../config/config.json"));
})[0];

export let options = {
    vus: config.k6.vus,
    duration: config.k6.duration,
};

export default function() {
    let res = http.get(config.k6.target_url);
    check(res, {
        "status is 200": (r) => r.status === 200,
        "response time < 500ms": (r) => r.timings.duration < 500
    });
    sleep(1);
}
