it('should handle arrays with Infinity and -Infinity', (done) => {
    const input = [1, Infinity, [-Infinity, [2]]];
    const expected = [1, Infinity, -Infinity, 2];
    const result = flattenArray(input);
    expect(result).toEqual(expected);
    done();
  });
  