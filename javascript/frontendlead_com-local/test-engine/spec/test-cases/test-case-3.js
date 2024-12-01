it('should flatten arrays containing various types including null and undefined', (done) => {
    const input = [null, undefined, [true, [false, ['string'], [{}]]]];
    const expected = [null, undefined, true, false, 'string', {}];
    const result = flattenArray(input);
    expect(result).toEqual(expected);
    done();
  });
  