import {describe, expect, it} from "@jest/globals";

describe('Hello', () => {
    it('should be Hello', () => {
        const name = "Daniel";
        expect(name).toEqual("Daniel");
    });
});
