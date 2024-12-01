it('should handle arrays with negative numbers', (done) => {
    const input = [1, -2, [3, -4, [-5]]];
    const expected = [1, -2, 3, -4, -5];
    const result = flattenArray(input);
    expect(result).toEqual(expected);
    done();
  });
  