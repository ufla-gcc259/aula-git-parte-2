using tempconv;
using Xunit;

namespace Tests
{
    [CollectionDefinition(nameof(tempconv))]
    public class TempcTest
    {
        public TempcTest()
        {
        }

        [Fact(DisplayName = "Teste - CToF")]
        public void Test_CToF()
        {
            var tempC = new TemperatureConverter();

            var result = tempC.CToF(Constants.absZeroCelsius);

            Assert.True(result.Equals(459));
        }

        [Fact(DisplayName = "Teste - FToC")]
        public void Test_FToC()
        {
            var tempC = new TemperatureConverter();

            var result = tempC.CToC(212);

            Assert.True(result.Equals(Constants.boilingCelsius));
        }
    }
}