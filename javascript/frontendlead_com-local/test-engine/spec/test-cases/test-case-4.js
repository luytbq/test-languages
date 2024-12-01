it('should return an empty array when the input is an empty array', (done) => {
    const input = [];
    const expected = [];
    const result = flattenArray(input);
    expect(result).toEqual(expected);
    done();
  });
  