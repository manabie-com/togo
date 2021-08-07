using Moq;
using NUnit.Framework;
using togo.Service;
using togo.Service.Dto;
using togo.Service.Errors;
using togo.Service.Interface;

namespace togo.UnitTest
{
    [TestFixture]
    public class TaskServiceTest
    {
        private ITaskService _taskService;

        [SetUp]
        public void SetUp()
        {
            var mockContext = TestHelper.GetMockTogoContext();
            var mockHttpContext = TestHelper.GetMockCurrentHttpContext();

            _taskService = new TaskService(mockContext, mockHttpContext);
        }

        [Test]
        [Order(1)]
        public void Create_Success()
        {
            var result = _taskService.Create(new TaskCreateDto { Content = "test" }).Result;
            Assert.AreEqual(result.Content, "test");
        }

        [Test]
        [Order(2)]
        public void Create_ReachRateLimit()
        {
            void _create()
            {
                try
                {
                    while (true)
                    {
                        _taskService.Create(new TaskCreateDto { Content = "test" }).Wait();
                    }
                }
                catch (System.Exception e)
                {

                    throw e.InnerException;
                }
            }

            Assert.Throws(Is.TypeOf<RestException>(), _create);
        }
    }
}
