import {sayHello} from "../src/say-hello";

describe(sayHello("Daniel"), () => {
    it(sayHello("Daniel"), () => {
        expect(sayHello("Daniel")).toEqual("Hello Daniel");
    })
});